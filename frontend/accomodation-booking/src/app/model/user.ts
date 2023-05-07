import { ObjectId } from 'bson';

export interface User {
  id: ObjectId | null;
  name: string;
  surname: string;
  email: string;
  password: string;
}
