import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { InsurancePayerService } from '../../../services/InsurancePayer.service';
import { InsurancePayer } from '../../../models/InsurancePayer';
import { SubBaseComponent } from '../../InsurancePayer/sub.base.component';

@Component({
    selector: 'app-create-insurancePayer',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateInsurancePayerComponent extends SubBaseComponent implements OnInit {

    title = 'Add InsurancePayer';

    insurancePayerForm: FormGroup;
    insurancePayer: InsurancePayer;

    constructor( http: HttpClient,
        private insurancePayerService: InsurancePayerService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.insurancePayerForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      website: ['', Validators.required],
      Plans: ['', ],
      Claims: ['', ],
      PayerType: ['', ]
        });
    }

    
    addInsurancePayer(name, website, Plans, Claims, PayerType): void {
        this.insurancePayerService
        .addInsurancePayer(name, website, Plans, Claims, PayerType)
            .subscribe(() => {
                this.router.navigate(['/indexInsurancePayer']);
            });
    }

    ngOnInit(): void {
    }
}