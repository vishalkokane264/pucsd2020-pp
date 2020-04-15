export class User {
   public id: number;
   public first_name: string;
   public last_name: string;
   public email:string;
   public contact_number:string;
   public updated_by:number;
   
   constructor(id: number,first_name: string,last_name: string,email:string,contact_number:string,updated_by:number)
   {
   this.id=id;
   this.first_name=first_name;
   this.last_name=last_name;
   this.email=email;
   this.contact_number=contact_number;
   this.updated_by=updated_by;
   }
} 
