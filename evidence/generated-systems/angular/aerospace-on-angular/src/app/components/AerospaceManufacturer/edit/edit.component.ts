import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { AerospaceManufacturerService } from '../../../services/AerospaceManufacturer.service';
import { SubBaseComponent } from '../../AerospaceManufacturer/sub.base.component';


@Component({
    selector: 'app-edit-aerospaceManufacturer',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditAerospaceManufacturerComponent extends SubBaseComponent implements OnInit {

    title = 'Edit AerospaceManufacturer';

    aerospaceManufacturerForm: FormGroup;
    aerospaceManufacturer: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: AerospaceManufacturerService,
        private fb: FormBuilder
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

    
    updateAerospaceManufacturer(name, legalName, headquartersCountry, website, Programs, Plants, Suppliers, ProductionCertificates): void {
        this.route.params.subscribe((params) => {

                        this.service.updateAerospaceManufacturer(name, legalName, headquartersCountry, website, Programs, Plants, Suppliers, ProductionCertificates, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexAerospaceManufacturer']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getAerospaceManufacturer(params['id']).subscribe(res => {
                this.aerospaceManufacturer = res;
            });
        });
    }
}