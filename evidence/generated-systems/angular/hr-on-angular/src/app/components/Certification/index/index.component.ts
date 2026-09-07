
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { CertificationService } from '../../../services/Certification.service';
import { Certification } from '../../../models/Certification';

@Component({
    selector: 'app-index-certification',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexCertificationComponent implements OnInit {

    certifications: Certification[] = [];

    constructor(
        private router: Router,
        private service: CertificationService
) {}

    ngOnInit(): void {
        this.getCertifications();
}

    getCertifications(): void {
        this.service.getCertifications().subscribe((res) => {
        this.certifications = res;
    });
}

    deleteCertification(id: any): void {
        this.service.deleteCertification(id)
            .subscribe(() => {
                this.getCertifications();
            });
    }
}