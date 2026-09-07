
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateCabinLayoutComponent } from './create.component';
import { CabinLayoutService } from '../../../services/CabinLayout.service';
import { Router } from '@angular/router';

describe('CreateCabinLayoutComponent', () => {
  let component: CreateCabinLayoutComponent;
  let fixture: ComponentFixture<CreateCabinLayoutComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateCabinLayoutComponent
      ],
      providers: [
        CabinLayoutService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateCabinLayoutComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});