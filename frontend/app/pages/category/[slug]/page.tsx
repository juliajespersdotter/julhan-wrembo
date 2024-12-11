import { Joke } from "@/app/components/joke";

export default async function Page({
  params,
}: {
  params: Promise<{ slug: string }>;
}) {
  const slug = (await params).slug;

  const data = await fetch("https://jjcc.wremert.work/category/" + slug);
  const jokes = await data?.json();

  return (
    <ul>
      {jokes &&
        jokes.map((joke: any) => (
          <>
            <Joke key={joke} joke={joke} />;
          </>
        ))}
    </ul>
  );
}
