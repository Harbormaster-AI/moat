import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { InsurerService } from '../../../services/Insurer.service';
import { Insurer } from '../../../models/Insurer';
import { SubBaseComponent } from '../../Insurer/sub.base.component';

@Component({
    selector: 'app-create-insurer',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateInsurerComponent extends SubBaseComponent implements OnInit {

    title = 'Add Insurer';

    insurerForm: FormGroup;
    insurer: Insurer;

    constructor( http: HttpClient,
        private insurerService: InsurerService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.insurerForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      legalName: ['', Validators.required],
      domicileCountry: ['', Validators.required],
      naicNumber: ['', Validators.required],
      website: ['', Validators.required],
      Products: ['', ],
      DistributionPartners: ['', ],
      Policies: ['', ],
      Claims: ['', ],
      ReinsuranceAgreements: ['', ]
        });
    }

    
    addInsurer(name, legalName, domicileCountry, naicNumber, website, Products, DistributionPartners, Policies, Claims, ReinsuranceAgreements): void {
        this.insurerService
        .addInsurer(name, legalName, domicileCountry, naicNumber, website, Products, DistributionPartners, Policies, Claims, ReinsuranceAgreements)
            .subscribe(() => {
                this.router.navigate(['/indexInsurer']);
            });
    }

    ngOnInit(): void {
    }
}