
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ClinicianService } from '../../../services/Clinician.service';
import { Clinician } from '../../../models/Clinician';

@Component({
    selector: 'app-index-clinician',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexClinicianComponent implements OnInit {

    clinicians: Clinician[] = [];

    constructor(
        private router: Router,
        private service: ClinicianService
) {}

    ngOnInit(): void {
        this.getClinicians();
}

    getClinicians(): void {
        this.service.getClinicians().subscribe((res) => {
        this.clinicians = res;
    });
}

    deleteClinician(id: any): void {
        this.service.deleteClinician(id)
            .subscribe(() => {
                this.getClinicians();
            });
    }
}