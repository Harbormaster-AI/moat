import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { TeamService } from '../../../services/Team.service';
import { Team } from '../../../models/Team';
import { SubBaseComponent } from '../../Team/sub.base.component';

@Component({
    selector: 'app-create-team',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateTeamComponent extends SubBaseComponent implements OnInit {

    title = 'Add Team';

    teamForm: FormGroup;
    team: Team;

    constructor( http: HttpClient,
        private teamService: TeamService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.teamForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      Agency: ['', ],
      Users: ['', ],
      AdAccounts: ['', ]
        });
    }

    
    addTeam(name, Agency, Users, AdAccounts): void {
        this.teamService
        .addTeam(name, Agency, Users, AdAccounts)
            .subscribe(() => {
                this.router.navigate(['/indexTeam']);
            });
    }

    ngOnInit(): void {
    }
}