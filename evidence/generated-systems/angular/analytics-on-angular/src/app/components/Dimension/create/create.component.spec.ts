
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateDimensionComponent } from './create.component';
import { DimensionService } from '../../../services/Dimension.service';
import { Router } from '@angular/router';

describe('CreateDimensionComponent', () => {
  let component: CreateDimensionComponent;
  let fixture: ComponentFixture<CreateDimensionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateDimensionComponent
      ],
      providers: [
        DimensionService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateDimensionComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});