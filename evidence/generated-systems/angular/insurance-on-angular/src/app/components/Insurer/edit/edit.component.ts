import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { InsurerService } from '../../../services/Insurer.service';
import { SubBaseComponent } from '../../Insurer/sub.base.component';


@Component({
    selector: 'app-edit-insurer',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditInsurerComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Insurer';

    insurerForm: FormGroup;
    insurer: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: InsurerService,
        private fb: FormBuilder
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

    
    updateInsurer(name, legalName, domicileCountry, naicNumber, website, Products, DistributionPartners, Policies, Claims, ReinsuranceAgreements): void {
        this.route.params.subscribe((params) => {

                        this.service.updateInsurer(name, legalName, domicileCountry, naicNumber, website, Products, DistributionPartners, Policies, Claims, ReinsuranceAgreements, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexInsurer']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getInsurer(params['id']).subscribe(res => {
                this.insurer = res;
            });
        });
    }
}