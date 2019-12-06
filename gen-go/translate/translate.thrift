namespace go translate

exception ServerError
{
    1: string code,
    2: string msg,
}

service Translate
{
    string translate(1:string src_string) throws(1:ServerError e);
}
