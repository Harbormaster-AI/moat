
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateSalesOrderComponent } from './create.component';
import { SalesOrderService } from '../../../services/SalesOrder.service';
import { Router } from '@angular/router';

describe('CreateSalesOrderComponent', () => {
  let component: CreateSalesOrderComponent;
  let fixture: ComponentFixture<CreateSalesOrderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateSalesOrderComponent
      ],
      providers: [
        SalesOrderService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateSalesOrderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});