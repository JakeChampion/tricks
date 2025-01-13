import { getStore } from "@netlify/blobs";

export default async (req, context) => {
    const construction = getStore("construction");
    await construction.set("nails", 9);
    const entry = await construction.get("nails")
    return new Response(entry);
};

export const config = {
    path: '/f'
};
