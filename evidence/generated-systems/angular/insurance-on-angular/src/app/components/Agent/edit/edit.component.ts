import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { AgentService } from '../../../services/Agent.service';
import { SubBaseComponent } from '../../Agent/sub.base.component';


@Component({
    selector: 'app-edit-agent',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditAgentComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Agent';

    agentForm: FormGroup;
    agent: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: AgentService,
        private fb: FormBuilder
) {
        super(http);
        this.agentForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  firstName: ['', Validators.required],
      lastName: ['', Validators.required],
      licenseId: ['', Validators.required],
      Distributor: ['', ],
      Policies: ['', ],
      Customers: ['', ],
      Status: ['', ]
        });
    }

    
    updateAgent(firstName, lastName, licenseId, Distributor, Policies, Customers, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateAgent(firstName, lastName, licenseId, Distributor, Policies, Customers, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexAgent']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getAgent(params['id']).subscribe(res => {
                this.agent = res;
            });
        });
    }
}