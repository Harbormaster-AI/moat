import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { RoutingService } from '../../../services/Routing.service';
import { SubBaseComponent } from '../../Routing/sub.base.component';


@Component({
    selector: 'app-edit-routing',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditRoutingComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Routing';

    routingForm: FormGroup;
    routing: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: RoutingService,
        private fb: FormBuilder
) {
        super(http);
        this.routingForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  routingNumber: ['', Validators.required],
      revision: ['', Validators.required],
      effectivityStart: ['', Validators.required],
      effectivityEnd: ['', Validators.required],
      Item: ['', ],
      Operations: ['', ],
      RoutingType: ['', ],
      Status: ['', ]
        });
    }

    
    updateRouting(routingNumber, revision, effectivityStart, effectivityEnd, Item, Operations, RoutingType, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateRouting(routingNumber, revision, effectivityStart, effectivityEnd, Item, Operations, RoutingType, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexRouting']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getRouting(params['id']).subscribe(res => {
                this.routing = res;
            });
        });
    }
}