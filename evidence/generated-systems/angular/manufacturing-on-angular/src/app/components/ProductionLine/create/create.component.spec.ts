
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateProductionLineComponent } from './create.component';
import { ProductionLineService } from '../../../services/ProductionLine.service';
import { Router } from '@angular/router';

describe('CreateProductionLineComponent', () => {
  let component: CreateProductionLineComponent;
  let fixture: ComponentFixture<CreateProductionLineComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateProductionLineComponent
      ],
      providers: [
        ProductionLineService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateProductionLineComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});