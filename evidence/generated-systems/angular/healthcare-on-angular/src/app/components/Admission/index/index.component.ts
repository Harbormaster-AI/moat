
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AdmissionService } from '../../../services/Admission.service';
import { Admission } from '../../../models/Admission';

@Component({
    selector: 'app-index-admission',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexAdmissionComponent implements OnInit {

    admissions: Admission[] = [];

    constructor(
        private router: Router,
        private service: AdmissionService
) {}

    ngOnInit(): void {
        this.getAdmissions();
}

    getAdmissions(): void {
        this.service.getAdmissions().subscribe((res) => {
        this.admissions = res;
    });
}

    deleteAdmission(id: any): void {
        this.service.deleteAdmission(id)
            .subscribe(() => {
                this.getAdmissions();
            });
    }
}