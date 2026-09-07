import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AerospaceManufacturerService } from '../../../services/AerospaceManufacturer.service';
import { AerospaceManufacturer } from '../../../models/AerospaceManufacturer';
import { SubBaseComponent } from '../../AerospaceManufacturer/sub.base.component';

@Component({
    selector: 'app-create-aerospaceManufacturer',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateAerospaceManufacturerComponent extends SubBaseComponent implements OnInit {

    title = 'Add AerospaceManufacturer';

    aerospaceManufacturerForm: FormGroup;
    aerospaceManufacturer: AerospaceManufacturer;

    constructor( http: HttpClient,
        private aerospaceManufacturerService: AerospaceManufacturerService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.aerospaceManufacturerForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      legalName: ['', Validators.required],
      headquartersCountry: ['', Validators.required],
      website: ['', Validators.required],
      Programs: ['', ],
      Plants: ['', ],
      Suppliers: ['', ],
      ProductionCertificates: ['', ]
        });
    }

    
    addAerospaceManufacturer(name, legalName, headquartersCountry, website, Programs, Plants, Suppliers, ProductionCertificates): void {
        this.aerospaceManufacturerService
        .addAerospaceManufacturer(name, legalName, headquartersCountry, website, Programs, Plants, Suppliers, ProductionCertificates)
            .subscribe(() => {
                this.router.navigate(['/indexAerospaceManufacturer']);
            });
    }

    ngOnInit(): void {
    }
}