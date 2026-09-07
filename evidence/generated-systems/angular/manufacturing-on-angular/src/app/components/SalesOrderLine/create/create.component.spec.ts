
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateSalesOrderLineComponent } from './create.component';
import { SalesOrderLineService } from '../../../services/SalesOrderLine.service';
import { Router } from '@angular/router';

describe('CreateSalesOrderLineComponent', () => {
  let component: CreateSalesOrderLineComponent;
  let fixture: ComponentFixture<CreateSalesOrderLineComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateSalesOrderLineComponent
      ],
      providers: [
        SalesOrderLineService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateSalesOrderLineComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});