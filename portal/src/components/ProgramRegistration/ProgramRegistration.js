import React, { useState, useEffect } from "react";
import styles from "./ProgramRegistration.module.css";
import axios from "axios";
import { useKeycloak } from "@react-keycloak/web";
import Form from "@rjsf/core";
import ListView from '../ListView/ListView';

function VaccineRegistration() {
    const { keycloak } = useKeycloak();
    const [formData, setFormData] = useState(null);
    const [programList, setProgramList] = useState([])

    useEffect(() => {
        getListOfRegisteredPrograms();
    },[]);

    const config = {
        headers: {
            Authorization: `Bearer ${keycloak.token} `,
            "Content-Type": "application/json",
        },
    };

    const schema = {
        type: "object",
        required: ["name", "description", "startDate"],
        properties: {
            programId: {
                type: "number",
                title: "Internal Program UID",
            },
            name: {
                type: "string",
                title: "Program Name",
            },
            description: {
                type: "string",
                title: "Program Description",
            },
            logoURL: {
                type: "string",
                title: "Program Logo",
                format: "file",
            },
            startDate: {
                type: "string",
                title: "Start Date",
                format: "date",
            },
            endDate: {
                type: "string",
                title: "End Date",
                format: "date",
            },
            status: {
                type: "string",
                title: "Status",
                enum: [
                    "Active", 
                    "Inactive"
                ],
            },
        },
    };

    

    const uiSchema = {
        classNames: styles["form"],
        title: {
            classNames: styles["form-title"],
        },
        status: {
            classNames: styles["form-radio-buttons"],
            "ui:widget": "radio",
            "ui:options": {
                "inline": true,
            }
        },
    };

    const handleSubmit = (datatoSend) => {
        axios
            .post("/divoc/admin/api/v1/programs", datatoSend, config)
            .then((res) => {
                alert("Status code is", res);
                console.log(res);
            });
    };

    const getListOfRegisteredPrograms = async () => {
        const res = await axios
            .get("/divoc/admin/api/v1/programs", config)
            .then( (res) => {
                return res.data
            })
        setProgramList(res)
    }

    return (
        <div className={styles["container"]}>
            <div className={styles["form-container"]}>
            <h4 className={styles['heading']}>Register New Vaccine</h4>
                <Form
                    schema={schema}
                    uiSchema={uiSchema}
                    onSubmit={(e) => {
                        setFormData(e.formData)
                        const newField = {medicineIds: [""]}
                        Object.assign(e.formData, newField)
                        handleSubmit(e.formData);
                    }}
                >
                    <button type="submit" className={styles['button']}>SAVE</button>
                </Form>
            </div>
            <div className={styles["sub-container"]}>
                <p className={styles['list-title']}>List of Registered Vaccine Programs</p>
                <ListView listData={programList} />
            </div>
        </div>
    );
}

export default VaccineRegistration;