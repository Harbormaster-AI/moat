
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { DocumentService } from '../../../services/Document.service';
import { Document } from '../../../models/Document';

@Component({
    selector: 'app-index-document',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexDocumentComponent implements OnInit {

    documents: Document[] = [];

    constructor(
        private router: Router,
        private service: DocumentService
) {}

    ngOnInit(): void {
        this.getDocuments();
}

    getDocuments(): void {
        this.service.getDocuments().subscribe((res) => {
        this.documents = res;
    });
}

    deleteDocument(id: any): void {
        this.service.deleteDocument(id)
            .subscribe(() => {
                this.getDocuments();
            });
    }
}