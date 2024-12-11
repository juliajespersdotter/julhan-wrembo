import React from "react";

export const Joke = ({ props }: any) => {
    return (
        <li key={props.id}>
            {props.telling} <br />
            {props.punchline}
            <br />
            <br />
        </li>
    );
};
