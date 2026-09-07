
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateInsurerComponent } from './create.component';
import { InsurerService } from '../../../services/Insurer.service';
import { Router } from '@angular/router';

describe('CreateInsurerComponent', () => {
  let component: CreateInsurerComponent;
  let fixture: ComponentFixture<CreateInsurerComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateInsurerComponent
      ],
      providers: [
        InsurerService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateInsurerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});