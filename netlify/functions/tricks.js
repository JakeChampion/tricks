import { getStore } from "@netlify/blobs";

async function f(request, init) {
    console.log({request, init})
    return fetch(request, init)
}

export default async (req, context) => {
    const construction = getStore({
        fetch: f, 
        name: "construction"
    });
    await construction.set("nails", 9);
    const entry = await construction.get("nails")
    return new Response(entry);
};
