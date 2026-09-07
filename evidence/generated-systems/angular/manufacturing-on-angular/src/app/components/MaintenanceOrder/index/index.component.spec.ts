
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexMaintenanceOrderComponent } from './index.component';
import { MaintenanceOrderService } from '../../../services/MaintenanceOrder.service';

describe('IndexMaintenanceOrderComponent', () => {
  let component: IndexMaintenanceOrderComponent;
  let fixture: ComponentFixture<IndexMaintenanceOrderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexMaintenanceOrderComponent
      ],
      providers: [
        MaintenanceOrderService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexMaintenanceOrderComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});