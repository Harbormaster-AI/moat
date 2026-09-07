import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { OrganizationService } from '../../../services/Organization.service';
import { SubBaseComponent } from '../../Organization/sub.base.component';


@Component({
    selector: 'app-edit-organization',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditOrganizationComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Organization';

    organizationForm: FormGroup;
    organization: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: OrganizationService,
        private fb: FormBuilder
) {
        super(http);
        this.organizationForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      legalName: ['', Validators.required],
      jurisdiction: ['', Validators.required],
      industrySector: ['', Validators.required],
      GovernanceBodies: ['', ],
      Policies: ['', ],
      Risks: ['', ],
      ThirdParties: ['', ],
      RecordsRepositories: ['', ],
      DataProcessingActivities: ['', ],
      CompliancePrograms: ['', ],
      AuditPrograms: ['', ],
      BusinessUnits: ['', ],
      Matters: ['', ],
      DataBreaches: ['', ]
        });
    }

    
    updateOrganization(name, legalName, jurisdiction, industrySector, GovernanceBodies, Policies, Risks, ThirdParties, RecordsRepositories, DataProcessingActivities, CompliancePrograms, AuditPrograms, BusinessUnits, Matters, DataBreaches): void {
        this.route.params.subscribe((params) => {

                        this.service.updateOrganization(name, legalName, jurisdiction, industrySector, GovernanceBodies, Policies, Risks, ThirdParties, RecordsRepositories, DataProcessingActivities, CompliancePrograms, AuditPrograms, BusinessUnits, Matters, DataBreaches, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexOrganization']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getOrganization(params['id']).subscribe(res => {
                this.organization = res;
            });
        });
    }
}