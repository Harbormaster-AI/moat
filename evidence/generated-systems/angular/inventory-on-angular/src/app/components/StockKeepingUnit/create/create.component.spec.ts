
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateStockKeepingUnitComponent } from './create.component';
import { StockKeepingUnitService } from '../../../services/StockKeepingUnit.service';
import { Router } from '@angular/router';

describe('CreateStockKeepingUnitComponent', () => {
  let component: CreateStockKeepingUnitComponent;
  let fixture: ComponentFixture<CreateStockKeepingUnitComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateStockKeepingUnitComponent
      ],
      providers: [
        StockKeepingUnitService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateStockKeepingUnitComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});