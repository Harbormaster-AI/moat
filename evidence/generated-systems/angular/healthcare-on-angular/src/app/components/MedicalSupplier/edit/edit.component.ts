import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { MedicalSupplierService } from '../../../services/MedicalSupplier.service';
import { SubBaseComponent } from '../../MedicalSupplier/sub.base.component';


@Component({
    selector: 'app-edit-medicalSupplier',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditMedicalSupplierComponent extends SubBaseComponent implements OnInit {

    title = 'Edit MedicalSupplier';

    medicalSupplierForm: FormGroup;
    medicalSupplier: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: MedicalSupplierService,
        private fb: FormBuilder
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

    
    updateMedicalSupplier(name, website, Facilities, InventoryItems, SupplierTier): void {
        this.route.params.subscribe((params) => {

                        this.service.updateMedicalSupplier(name, website, Facilities, InventoryItems, SupplierTier, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexMedicalSupplier']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getMedicalSupplier(params['id']).subscribe(res => {
                this.medicalSupplier = res;
            });
        });
    }
}