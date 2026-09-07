import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { PharmacyService } from '../../../services/Pharmacy.service';
import { SubBaseComponent } from '../../Pharmacy/sub.base.component';


@Component({
    selector: 'app-edit-pharmacy',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditPharmacyComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Pharmacy';

    pharmacyForm: FormGroup;
    pharmacy: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: PharmacyService,
        private fb: FormBuilder
) {
        super(http);
        this.pharmacyForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      Facility: ['', ],
      MedicationDispenses: ['', ],
      MedicationOrders: ['', ]
        });
    }

    
    updatePharmacy(name, Facility, MedicationDispenses, MedicationOrders): void {
        this.route.params.subscribe((params) => {

                        this.service.updatePharmacy(name, Facility, MedicationDispenses, MedicationOrders, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexPharmacy']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getPharmacy(params['id']).subscribe(res => {
                this.pharmacy = res;
            });
        });
    }
}