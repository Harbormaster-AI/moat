import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { TerritoryService } from '../../../services/Territory.service';
import { Territory } from '../../../models/Territory';
import { SubBaseComponent } from '../../Territory/sub.base.component';

@Component({
    selector: 'app-create-territory',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateTerritoryComponent extends SubBaseComponent implements OnInit {

    title = 'Add Territory';

    territoryForm: FormGroup;
    territory: Territory;

    constructor( http: HttpClient,
        private territoryService: TerritoryService,
        private fb: FormBuilder,
        private router: Router
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

    
    addTerritory(name, region, Organization, Accounts, Users, TerritoryType): void {
        this.territoryService
        .addTerritory(name, region, Organization, Accounts, Users, TerritoryType)
            .subscribe(() => {
                this.router.navigate(['/indexTerritory']);
            });
    }

    ngOnInit(): void {
    }
}