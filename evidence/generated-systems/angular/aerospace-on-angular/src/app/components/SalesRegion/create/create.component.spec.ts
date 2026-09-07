
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateSalesRegionComponent } from './create.component';
import { SalesRegionService } from '../../../services/SalesRegion.service';
import { Router } from '@angular/router';

describe('CreateSalesRegionComponent', () => {
  let component: CreateSalesRegionComponent;
  let fixture: ComponentFixture<CreateSalesRegionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateSalesRegionComponent
      ],
      providers: [
        SalesRegionService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateSalesRegionComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});