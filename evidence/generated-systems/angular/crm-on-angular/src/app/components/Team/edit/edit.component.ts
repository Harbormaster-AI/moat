import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { TeamService } from '../../../services/Team.service';
import { SubBaseComponent } from '../../Team/sub.base.component';


@Component({
    selector: 'app-edit-team',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditTeamComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Team';

    teamForm: FormGroup;
    team: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: TeamService,
        private fb: FormBuilder
) {
        super(http);
        this.teamForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      Organization: ['', ],
      Users: ['', ],
      Accounts: ['', ],
      Opportunities: ['', ],
      Cases: ['', ],
      Campaigns: ['', ],
      TeamType: ['', ]
        });
    }

    
    updateTeam(name, Organization, Users, Accounts, Opportunities, Cases, Campaigns, TeamType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateTeam(name, Organization, Users, Accounts, Opportunities, Cases, Campaigns, TeamType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexTeam']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getTeam(params['id']).subscribe(res => {
                this.team = res;
            });
        });
    }
}