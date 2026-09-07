
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { MedicationDispenseService } from '../../../services/MedicationDispense.service';
import { MedicationDispense } from '../../../models/MedicationDispense';

@Component({
    selector: 'app-index-medicationDispense',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexMedicationDispenseComponent implements OnInit {

    medicationDispenses: MedicationDispense[] = [];

    constructor(
        private router: Router,
        private service: MedicationDispenseService
) {}

    ngOnInit(): void {
        this.getMedicationDispenses();
}

    getMedicationDispenses(): void {
        this.service.getMedicationDispenses().subscribe((res) => {
        this.medicationDispenses = res;
    });
}

    deleteMedicationDispense(id: any): void {
        this.service.deleteMedicationDispense(id)
            .subscribe(() => {
                this.getMedicationDispenses();
            });
    }
}