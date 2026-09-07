
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateUoMConversionComponent } from './create.component';
import { UoMConversionService } from '../../../services/UoMConversion.service';
import { Router } from '@angular/router';

describe('CreateUoMConversionComponent', () => {
  let component: CreateUoMConversionComponent;
  let fixture: ComponentFixture<CreateUoMConversionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateUoMConversionComponent
      ],
      providers: [
        UoMConversionService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateUoMConversionComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});