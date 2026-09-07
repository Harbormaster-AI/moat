
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexWorkOrderComponent } from './index.component';
import { WorkOrderService } from '../../../services/WorkOrder.service';

describe('IndexWorkOrderComponent', () => {
  let component: IndexWorkOrderComponent;
  let fixture: ComponentFixture<IndexWorkOrderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexWorkOrderComponent
      ],
      providers: [
        WorkOrderService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexWorkOrderComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});