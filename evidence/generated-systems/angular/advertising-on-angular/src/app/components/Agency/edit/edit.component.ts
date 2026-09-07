import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { AgencyService } from '../../../services/Agency.service';
import { SubBaseComponent } from '../../Agency/sub.base.component';


@Component({
    selector: 'app-edit-agency',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditAgencyComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Agency';

    agencyForm: FormGroup;
    agency: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: AgencyService,
        private fb: FormBuilder
) {
        super(http);
        this.agencyForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      legalName: ['', Validators.required],
      headquartersCountry: ['', Validators.required],
      website: ['', Validators.required],
      Advertisers: ['', ],
      Teams: ['', ],
      Users: ['', ],
      InsertionOrders: ['', ]
        });
    }

    
    updateAgency(name, legalName, headquartersCountry, website, Advertisers, Teams, Users, InsertionOrders): void {
        this.route.params.subscribe((params) => {

                        this.service.updateAgency(name, legalName, headquartersCountry, website, Advertisers, Teams, Users, InsertionOrders, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexAgency']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getAgency(params['id']).subscribe(res => {
                this.agency = res;
            });
        });
    }
}