import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AgentService } from '../../../services/Agent.service';
import { Agent } from '../../../models/Agent';
import { SubBaseComponent } from '../../Agent/sub.base.component';

@Component({
    selector: 'app-create-agent',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateAgentComponent extends SubBaseComponent implements OnInit {

    title = 'Add Agent';

    agentForm: FormGroup;
    agent: Agent;

    constructor( http: HttpClient,
        private agentService: AgentService,
        private fb: FormBuilder,
        private router: Router
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

    
    addAgent(firstName, lastName, licenseId, Distributor, Policies, Customers, Status): void {
        this.agentService
        .addAgent(firstName, lastName, licenseId, Distributor, Policies, Customers, Status)
            .subscribe(() => {
                this.router.navigate(['/indexAgent']);
            });
    }

    ngOnInit(): void {
    }
}