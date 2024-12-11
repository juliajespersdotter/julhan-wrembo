"use client";

//Get all categories in a list, link to slug page with category name

export default async function Category() {
    const data = await fetch(process.env.NEXT_PUBLIC_BASE_API_URL + "/categories/");
    const categories = await data?.json();

    return (
        <ul>
            {categories &&
                categories.map((category: any) => (
                    <>
                        <br />
                        <li key={category}>
                            <a href={`/pages/category/${category.name}`}> {category.name}</a>
                        </li>
                    </>
                ))}
        </ul>
    );
}
