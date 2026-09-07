
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexMaintenancePlanComponent } from './index.component';
import { MaintenancePlanService } from '../../../services/MaintenancePlan.service';

describe('IndexMaintenancePlanComponent', () => {
  let component: IndexMaintenancePlanComponent;
  let fixture: ComponentFixture<IndexMaintenancePlanComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexMaintenancePlanComponent
      ],
      providers: [
        MaintenancePlanService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexMaintenancePlanComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});