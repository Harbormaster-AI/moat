
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { TypeCertificateService } from '../../../services/TypeCertificate.service';
import { TypeCertificate } from '../../../models/TypeCertificate';

@Component({
    selector: 'app-index-typeCertificate',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexTypeCertificateComponent implements OnInit {

    typeCertificates: TypeCertificate[] = [];

    constructor(
        private router: Router,
        private service: TypeCertificateService
) {}

    ngOnInit(): void {
        this.getTypeCertificates();
}

    getTypeCertificates(): void {
        this.service.getTypeCertificates().subscribe((res) => {
        this.typeCertificates = res;
    });
}

    deleteTypeCertificate(id: any): void {
        this.service.deleteTypeCertificate(id)
            .subscribe(() => {
                this.getTypeCertificates();
            });
    }
}