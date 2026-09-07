import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { PharmacyService } from '../../../services/Pharmacy.service';
import { Pharmacy } from '../../../models/Pharmacy';
import { SubBaseComponent } from '../../Pharmacy/sub.base.component';

@Component({
    selector: 'app-create-pharmacy',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreatePharmacyComponent extends SubBaseComponent implements OnInit {

    title = 'Add Pharmacy';

    pharmacyForm: FormGroup;
    pharmacy: Pharmacy;

    constructor( http: HttpClient,
        private pharmacyService: PharmacyService,
        private fb: FormBuilder,
        private router: Router
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

    
    addPharmacy(name, Facility, MedicationDispenses, MedicationOrders): void {
        this.pharmacyService
        .addPharmacy(name, Facility, MedicationDispenses, MedicationOrders)
            .subscribe(() => {
                this.router.navigate(['/indexPharmacy']);
            });
    }

    ngOnInit(): void {
    }
}