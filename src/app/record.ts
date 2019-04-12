import { Metadata } from './metadata';

export class Record {
    file: Blob;
    filename: string;
    id: string;
    metadata: Metadata;
}
