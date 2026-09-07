import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { AnalyticsWorkspaceService } from '../../../services/AnalyticsWorkspace.service';
import { SubBaseComponent } from '../../AnalyticsWorkspace/sub.base.component';


@Component({
    selector: 'app-edit-analyticsWorkspace',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditAnalyticsWorkspaceComponent extends SubBaseComponent implements OnInit {

    title = 'Edit AnalyticsWorkspace';

    analyticsWorkspaceForm: FormGroup;
    analyticsWorkspace: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: AnalyticsWorkspaceService,
        private fb: FormBuilder
) {
        super(http);
        this.analyticsWorkspaceForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      businessDomain: ['', Validators.required],
      ownerTeam: ['', Validators.required],
      Datasets: ['', ],
      DataSources: ['', ],
      Pipelines: ['', ],
      Dashboards: ['', ],
      Reports: ['', ],
      Notebooks: ['', ],
      Models: ['', ],
      FeatureSets: ['', ],
      Policies: ['', ],
      LineageNodes: ['', ],
      GovernanceTier: ['', ]
        });
    }

    
    updateAnalyticsWorkspace(name, businessDomain, ownerTeam, Datasets, DataSources, Pipelines, Dashboards, Reports, Notebooks, Models, FeatureSets, Policies, LineageNodes, GovernanceTier): void {
        this.route.params.subscribe((params) => {

                        this.service.updateAnalyticsWorkspace(name, businessDomain, ownerTeam, Datasets, DataSources, Pipelines, Dashboards, Reports, Notebooks, Models, FeatureSets, Policies, LineageNodes, GovernanceTier, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexAnalyticsWorkspace']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getAnalyticsWorkspace(params['id']).subscribe(res => {
                this.analyticsWorkspace = res;
            });
        });
    }
}