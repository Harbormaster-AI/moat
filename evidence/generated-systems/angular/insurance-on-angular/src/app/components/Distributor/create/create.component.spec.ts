
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateDistributorComponent } from './create.component';
import { DistributorService } from '../../../services/Distributor.service';
import { Router } from '@angular/router';

describe('CreateDistributorComponent', () => {
  let component: CreateDistributorComponent;
  let fixture: ComponentFixture<CreateDistributorComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateDistributorComponent
      ],
      providers: [
        DistributorService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateDistributorComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});