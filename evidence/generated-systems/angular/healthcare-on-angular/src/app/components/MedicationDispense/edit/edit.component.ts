import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { MedicationDispenseService } from '../../../services/MedicationDispense.service';
import { SubBaseComponent } from '../../MedicationDispense/sub.base.component';


@Component({
    selector: 'app-edit-medicationDispense',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditMedicationDispenseComponent extends SubBaseComponent implements OnInit {

    title = 'Edit MedicationDispense';

    medicationDispenseForm: FormGroup;
    medicationDispense: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: MedicationDispenseService,
        private fb: FormBuilder
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

    
    updateMedicationDispense(dispenseNumber, quantity, whenPrepared, MedicationOrder, Pharmacy, Patient, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateMedicationDispense(dispenseNumber, quantity, whenPrepared, MedicationOrder, Pharmacy, Patient, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexMedicationDispense']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getMedicationDispense(params['id']).subscribe(res => {
                this.medicationDispense = res;
            });
        });
    }
}