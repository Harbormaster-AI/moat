
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { KYCDocumentService } from '../../../services/KYCDocument.service';
import { KYCDocument } from '../../../models/KYCDocument';

@Component({
    selector: 'app-index-kYCDocument',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexKYCDocumentComponent implements OnInit {

    kYCDocuments: KYCDocument[] = [];

    constructor(
        private router: Router,
        private service: KYCDocumentService
) {}

    ngOnInit(): void {
        this.getKYCDocuments();
}

    getKYCDocuments(): void {
        this.service.getKYCDocuments().subscribe((res) => {
        this.kYCDocuments = res;
    });
}

    deleteKYCDocument(id: any): void {
        this.service.deleteKYCDocument(id)
            .subscribe(() => {
                this.getKYCDocuments();
            });
    }
}