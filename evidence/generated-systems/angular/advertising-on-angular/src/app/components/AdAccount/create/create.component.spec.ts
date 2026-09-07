
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateAdAccountComponent } from './create.component';
import { AdAccountService } from '../../../services/AdAccount.service';
import { Router } from '@angular/router';

describe('CreateAdAccountComponent', () => {
  let component: CreateAdAccountComponent;
  let fixture: ComponentFixture<CreateAdAccountComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateAdAccountComponent
      ],
      providers: [
        AdAccountService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateAdAccountComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});