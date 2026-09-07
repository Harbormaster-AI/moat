import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AgencyService } from '../../../services/Agency.service';
import { Agency } from '../../../models/Agency';
import { SubBaseComponent } from '../../Agency/sub.base.component';

@Component({
    selector: 'app-create-agency',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateAgencyComponent extends SubBaseComponent implements OnInit {

    title = 'Add Agency';

    agencyForm: FormGroup;
    agency: Agency;

    constructor( http: HttpClient,
        private agencyService: AgencyService,
        private fb: FormBuilder,
        private router: Router
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

    
    addAgency(name, legalName, headquartersCountry, website, Advertisers, Teams, Users, InsertionOrders): void {
        this.agencyService
        .addAgency(name, legalName, headquartersCountry, website, Advertisers, Teams, Users, InsertionOrders)
            .subscribe(() => {
                this.router.navigate(['/indexAgency']);
            });
    }

    ngOnInit(): void {
    }
}