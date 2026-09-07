import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ProductionCertificateService } from '../../../services/ProductionCertificate.service';
import { ProductionCertificate } from '../../../models/ProductionCertificate';
import { SubBaseComponent } from '../../ProductionCertificate/sub.base.component';

@Component({
    selector: 'app-create-productionCertificate',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateProductionCertificateComponent extends SubBaseComponent implements OnInit {

    title = 'Add ProductionCertificate';

    productionCertificateForm: FormGroup;
    productionCertificate: ProductionCertificate;

    constructor( http: HttpClient,
        private productionCertificateService: ProductionCertificateService,
        private fb: FormBuilder,
        private router: Router
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

    
    addProductionCertificate(certificateNumber, authority, Manufacturer): void {
        this.productionCertificateService
        .addProductionCertificate(certificateNumber, authority, Manufacturer)
            .subscribe(() => {
                this.router.navigate(['/indexProductionCertificate']);
            });
    }

    ngOnInit(): void {
    }
}