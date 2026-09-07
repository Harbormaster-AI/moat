
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateTerritoryComponent } from './create.component';
import { TerritoryService } from '../../../services/Territory.service';
import { Router } from '@angular/router';

describe('CreateTerritoryComponent', () => {
  let component: CreateTerritoryComponent;
  let fixture: ComponentFixture<CreateTerritoryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateTerritoryComponent
      ],
      providers: [
        TerritoryService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateTerritoryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});