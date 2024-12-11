import { Joke } from "@/app/components/joke";
import React from "react";

export default async function Page() {
    let data = await fetch(process.env.NEXT_PUBLIC_BASE_API_URL + "/jokes");
    let jokes = await data?.json();

    return (
        <ul>
            {jokes &&
                jokes.map((joke: any) => (
                    <>
                        <Joke key={joke.id} props={joke} />
                    </>
                ))}
        </ul>
    );
}
