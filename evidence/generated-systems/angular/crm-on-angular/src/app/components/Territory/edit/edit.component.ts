import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { TerritoryService } from '../../../services/Territory.service';
import { SubBaseComponent } from '../../Territory/sub.base.component';


@Component({
    selector: 'app-edit-territory',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditTerritoryComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Territory';

    territoryForm: FormGroup;
    territory: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: TerritoryService,
        private fb: FormBuilder
) {
        super(http);
        this.territoryForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      region: ['', Validators.required],
      Organization: ['', ],
      Accounts: ['', ],
      Users: ['', ],
      TerritoryType: ['', ]
        });
    }

    
    updateTerritory(name, region, Organization, Accounts, Users, TerritoryType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateTerritory(name, region, Organization, Accounts, Users, TerritoryType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexTerritory']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getTerritory(params['id']).subscribe(res => {
                this.territory = res;
            });
        });
    }
}