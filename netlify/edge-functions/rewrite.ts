import { getStore } from "@netlify/blobs";
import type { Context } from "@netlify/functions";

let r = [];
let i = [];
async function f(request, init) {
    r.push(request) 
    i.push(init)
    console.log({request, init})
    return fetch(request, init)
}

export default async (req: Request, context: Context) => {
    const construction = getStore({ fetch: f, name: "construction" });
    await construction.set("nails", 9);
    const entry = await construction.get("nails")
    return Response.json({r,i,entry});
};

export const config = {
    path: '/*'
};
