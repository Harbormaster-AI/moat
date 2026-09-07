import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MedicationDispenseService } from '../../../services/MedicationDispense.service';
import { MedicationDispense } from '../../../models/MedicationDispense';
import { SubBaseComponent } from '../../MedicationDispense/sub.base.component';

@Component({
    selector: 'app-create-medicationDispense',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateMedicationDispenseComponent extends SubBaseComponent implements OnInit {

    title = 'Add MedicationDispense';

    medicationDispenseForm: FormGroup;
    medicationDispense: MedicationDispense;

    constructor( http: HttpClient,
        private medicationDispenseService: MedicationDispenseService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.medicationDispenseForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  dispenseNumber: ['', Validators.required],
      quantity: ['', Validators.required],
      whenPrepared: ['', Validators.required],
      MedicationOrder: ['', ],
      Pharmacy: ['', ],
      Patient: ['', ],
      Status: ['', ]
        });
    }

    
    addMedicationDispense(dispenseNumber, quantity, whenPrepared, MedicationOrder, Pharmacy, Patient, Status): void {
        this.medicationDispenseService
        .addMedicationDispense(dispenseNumber, quantity, whenPrepared, MedicationOrder, Pharmacy, Patient, Status)
            .subscribe(() => {
                this.router.navigate(['/indexMedicationDispense']);
            });
    }

    ngOnInit(): void {
    }
}