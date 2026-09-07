import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MedicalSupplierService } from '../../../services/MedicalSupplier.service';
import { MedicalSupplier } from '../../../models/MedicalSupplier';
import { SubBaseComponent } from '../../MedicalSupplier/sub.base.component';

@Component({
    selector: 'app-create-medicalSupplier',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateMedicalSupplierComponent extends SubBaseComponent implements OnInit {

    title = 'Add MedicalSupplier';

    medicalSupplierForm: FormGroup;
    medicalSupplier: MedicalSupplier;

    constructor( http: HttpClient,
        private medicalSupplierService: MedicalSupplierService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.medicalSupplierForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      website: ['', Validators.required],
      Facilities: ['', ],
      InventoryItems: ['', ],
      SupplierTier: ['', ]
        });
    }

    
    addMedicalSupplier(name, website, Facilities, InventoryItems, SupplierTier): void {
        this.medicalSupplierService
        .addMedicalSupplier(name, website, Facilities, InventoryItems, SupplierTier)
            .subscribe(() => {
                this.router.navigate(['/indexMedicalSupplier']);
            });
    }

    ngOnInit(): void {
    }
}