import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ProductionCertificateService } from '../../../services/ProductionCertificate.service';
import { SubBaseComponent } from '../../ProductionCertificate/sub.base.component';


@Component({
    selector: 'app-edit-productionCertificate',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditProductionCertificateComponent extends SubBaseComponent implements OnInit {

    title = 'Edit ProductionCertificate';

    productionCertificateForm: FormGroup;
    productionCertificate: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ProductionCertificateService,
        private fb: FormBuilder
) {
        super(http);
        this.productionCertificateForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  certificateNumber: ['', Validators.required],
      authority: ['', Validators.required],
      Manufacturer: ['', ]
        });
    }

    
    updateProductionCertificate(certificateNumber, authority, Manufacturer): void {
        this.route.params.subscribe((params) => {

                        this.service.updateProductionCertificate(certificateNumber, authority, Manufacturer, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexProductionCertificate']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getProductionCertificate(params['id']).subscribe(res => {
                this.productionCertificate = res;
            });
        });
    }
}