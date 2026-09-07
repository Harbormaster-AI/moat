
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { PatientService } from '../../../services/Patient.service';
import { Patient } from '../../../models/Patient';

@Component({
    selector: 'app-index-patient',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexPatientComponent implements OnInit {

    patients: Patient[] = [];

    constructor(
        private router: Router,
        private service: PatientService
) {}

    ngOnInit(): void {
        this.getPatients();
}

    getPatients(): void {
        this.service.getPatients().subscribe((res) => {
        this.patients = res;
    });
}

    deletePatient(id: any): void {
        this.service.deletePatient(id)
            .subscribe(() => {
                this.getPatients();
            });
    }
}