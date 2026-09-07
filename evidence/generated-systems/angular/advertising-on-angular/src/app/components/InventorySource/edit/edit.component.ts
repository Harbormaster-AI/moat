import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { InventorySourceService } from '../../../services/InventorySource.service';
import { SubBaseComponent } from '../../InventorySource/sub.base.component';


@Component({
    selector: 'app-edit-inventorySource',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditInventorySourceComponent extends SubBaseComponent implements OnInit {

    title = 'Edit InventorySource';

    inventorySourceForm: FormGroup;
    inventorySource: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: InventorySourceService,
        private fb: FormBuilder
) {
        super(http);
        this.inventorySourceForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      domain: ['', Validators.required],
      Publisher: ['', ],
      AdSlots: ['', ],
      Deals: ['', ],
      Channel: ['', ],
      PrimaryFormat: ['', ]
        });
    }

    
    updateInventorySource(name, domain, Publisher, AdSlots, Deals, Channel, PrimaryFormat): void {
        this.route.params.subscribe((params) => {

                        this.service.updateInventorySource(name, domain, Publisher, AdSlots, Deals, Channel, PrimaryFormat, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexInventorySource']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getInventorySource(params['id']).subscribe(res => {
                this.inventorySource = res;
            });
        });
    }
}